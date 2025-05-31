using SecureChat.src.db.repository.UserRepository;
using SecureChat.src.db.schema;
using SecureChat.src.utils;

namespace SecureChat.src.service.UserService
{
    public class UserService(IUserRepository userRepository) : IUserService
    {
        private readonly IUserRepository _userRepository = userRepository;

        async Task<User> IUserService.RegisterAsync(User user)
        {
            if (!string.IsNullOrEmpty(user.Email))
            {
                var existingUser = await _userRepository.GetUserByEmailAsync(user.Email);
                if (existingUser != null)
                {
                    throw new InvalidOperationException("Email already in use");
                }
            }
            if (!string.IsNullOrEmpty(user.PhoneNumber))
            {
                var existingUser = await _userRepository.GetUserByPhoneNumberAsync(user.PhoneNumber);
                if (existingUser != null)
                {
                    throw new InvalidOperationException("Phone number already in use");
                }
            }
            user.Password = HashUtils.HashPassword(user.Password);
            user.Password = HashUtils.HashPassword(user.Password);
            user.UpdatedAt = DateTime.UtcNow;
            await _userRepository.AddUserAsync(user);
            return user;
        }

        public Task<User> RegisterAsync(User user)
        {
            throw new NotImplementedException();
        }

        public Task<User> LoginEmailAsync(string email, string password)
        {
            throw new NotImplementedException();
        }

        public Task<User> LoginPhoneNumberAsync(string phoneNumber, string password)
        {
            throw new NotImplementedException();
        }

        public Task<User> GetUserByIdAsync(int id)
        {
            throw new NotImplementedException();
        }

        public Task<User> GetUserByEmailAsync(string email)
        {
            throw new NotImplementedException();
        }

        public Task<User> GetUserByPhoneNumberAsync(string phoneNumber)
        {
            throw new NotImplementedException();
        }
    }
}