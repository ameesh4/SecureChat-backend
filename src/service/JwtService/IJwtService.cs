using SecureChat.db.schema;

namespace SecureChat.service.JwtService
{
    public interface IJwtService
    {
        string GenerateToken(string userId);
        Task<User?> ValidateToken(string token);

        string GenerateRefreshToken(string userId);
        Task<User?> ValidateRefreshToken(string refreshToken);
    }
}