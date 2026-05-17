package request

func ParseCreateUser(c fiber.Ctx) (user.UserCreateInput, error) {
	var req user.ReqCreate
	if err := c.Bind().Body(&req); err != nil {
		return user.UserCreateInput{}, err
	}

	return req.NewUser, nil
}

func ParseUpdateUser(c fiber.Ctx) (user.UserUpdate, error) {
	var req user.ReqUpdate
	if err := c.Bind().Body(&req); err != nil {
		return user.UserUpdate{}, err
	}

	return req.ToUpdate, nil
}
