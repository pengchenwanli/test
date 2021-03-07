package handler

func (h *Handler) initRouter() {
	api := h.engine.Group("") //路由分组
	api.POST("admin/new", h.NewAdmin)
	api.POST("admin/login", h.Login)
	{
		auth := api.Group("", h.SessionVerifier)
		{
			admin := auth.Group("admin")
			admin.POST("logout", h.Logout)
		}
		{

		}

	}
}
