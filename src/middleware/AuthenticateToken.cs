


using SecureChat.service.JwtService;

namespace SecureChat.api.middleware
{
    public class AuthenticateToken
    {
        void Invoke(HttpContext context, RequestDelegate next)
        {
            var token = context.Request.Headers["Authorization"].FirstOrDefault()?.Split(" ").Last();
            if (string.IsNullOrEmpty(token))
            {
                context.Response.StatusCode = StatusCodes.Status401Unauthorized;
                return;
            }

            var jwtService = context.RequestServices.GetRequiredService<IJwtService>();
            var user = jwtService.ValidateToken(token).Result;

            if (user == null)
            {
                context.Response.StatusCode = StatusCodes.Status401Unauthorized;
                return;
            }

            context.Items["User"] = user;
            next(context);
        }
    }
}