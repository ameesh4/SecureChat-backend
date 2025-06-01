using SecureChat.src.db.schema;
using SecureChat.src.api.model;


namespace SecureChat.src.service.UserService
{
    public interface IUserService
    {
        Task<bool> RegisterAsync(User user);
        Task<LoginResponse> LoginEmailAsync(UserLogin user);
        Task<User> LoginPhoneNumberAsync(string phoneNumber, string password);
        Task<User> GetUserByIdAsync(int id);
        Task<User> GetUserByEmailAsync(string email);
        Task<User> GetUserByPhoneNumberAsync(string phoneNumber);
    }
}