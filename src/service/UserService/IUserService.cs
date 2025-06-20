using SecureChat.api.model;
using SecureChat.db.schema;

namespace SecureChat.service.UserService
{
    public interface IUserService
    {
        Task<bool> RegisterAsync(Register user);
        Task<LoginResponse> LoginEmailAsync(UserLoginEmail user);
        Task<LoginResponse> LoginPhoneNumberAsync(UserLoginPhone user);
        Task<User> GetUserByIdAsync(int id);
        Task<User> GetUserByEmailAsync(string email);
        Task<User> GetUserByPhoneNumberAsync(string phoneNumber);
    }
}