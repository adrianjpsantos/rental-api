package application

import (
	"context"
	"fmt"

	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/profile"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
	"github.com/adrianjpsantos/rental-api/internal/security"
	"github.com/adrianjpsantos/rental-api/internal/uow"
)

type AuthService struct {
	UOW            uow.UnitOfWork
	SessionService session.Service
}

func (a *AuthService) Register(
	ctx context.Context,
	input authenticate.RegisterInput,
) error {
	fmt.Println("RegisterInput: ", input.Email)
	return a.UOW.Do(ctx, true, func(repositories repositories.AllRepositories) error {

		// 1. Criar User
		userID, err := repositories.User.Create(ctx, user.RoleUser)
		if err != nil {
			fmt.Println("REGISTER, Create User:", err)
			return err
		}

		// 2. Criar Profile
		profileInput := profile.CreateInput{
			UserID:    *userID,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			CPF:       input.CPF,
			Phone:     input.Phone,
			BirthDate: input.BirthDate,
		}

		err = repositories.Profile.Create(ctx, &profileInput)
		if err != nil {
			fmt.Println("REGISTER, Create Profile:", err)
			return err
		}

		// 3. Hash da senha
		passwordHash, err := security.GenerateHashedPassword(input.Password)
		if err != nil {
			return err
		}

		// 4. Criar AuthAccount
		accountInput := authaccount.CreateInput{
			UserID:         *userID,
			Provider:       authaccount.ProviderLocal,
			ProviderUserID: input.Email,
			Email:          input.Email,
			PasswordHash:   &passwordHash,
			IsPrimary:      true,
		}

		fmt.Println("AccountInput:", accountInput.Email)

		_, err = repositories.AuthAccount.Create(ctx, &accountInput)
		if err != nil {
			fmt.Println("REGISTER, Create Account:", err)
			return err
		}

		return nil
	})
}

func (a *AuthService) Logout(ctx context.Context, userID string) error {
	panic("unimplemented")
}

func (a *AuthService) LoginLocal(
	ctx context.Context,
	input authenticate.AuthenticateInput,
) (*authenticate.AuthenticateOutput, error) {

	var output *authenticate.AuthenticateOutput

	err := a.UOW.Do(ctx, false, func(repositories repositories.AllRepositories) error {

		authAccount, err := repositories.AuthAccount.FindLocalByEmail(ctx, input.Email)

		if err != nil {
			fmt.Println("Erro ao procurar email")
			return authenticate.ErrInvalidCredentials
		}

		err = security.CheckPassword(
			*authAccount.PasswordHash,
			input.Password,
		)

		if err != nil {
			fmt.Println("Erro ao comparar senha")
			return authenticate.ErrInvalidCredentials
		}
		refreshToken, accessToken, err := a.SessionService.StartSession(ctx, authAccount.ID)

		if err != nil {
			fmt.Println("Authenticate: ", err)
			return err

		}

		output = &authenticate.AuthenticateOutput{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}

func (a *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	return a.SessionService.RefreshSession(ctx, refreshToken)
}

func NewAuthService(uow uow.UnitOfWork, sessionService session.Service) authenticate.Service {
	return &AuthService{
		UOW:            uow,
		SessionService: sessionService,
	}
}
