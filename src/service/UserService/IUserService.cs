using hushline.src.db.schema;


namespace hushline.src.service.UserService
{
    public interface IUserService
    {
        Task<User> RegisterAsync(User user);
        Task<User> LoginEmailAsync(string email, string password);
        Task<User> LoginPhoneNumberAsync(string phoneNumber, string password);
        Task<User> GetUserByIdAsync(int id);
        Task<User> GetUserByEmailAsync(string email);
        Task<User> GetUserByPhoneNumberAsync(string phoneNumber);
    }
}