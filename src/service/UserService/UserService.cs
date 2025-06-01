using SecureChat.src.db.repository.UserRepository;
using SecureChat.src.db.schema;
using SecureChat.src.utils;
using SecureChat.src.api.model;
using SecureChat.src.service.JwtService;

namespace SecureChat.src.service.UserService
{
    public class UserService(IUserRepository userRepository, IJwtService jwtService) : IUserService
    {
        private readonly IUserRepository _userRepository = userRepository;
        private readonly IJwtService _jwtUtils = jwtService;

        async Task<bool> IUserService.RegisterAsync(User user)
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

            if (string.IsNullOrEmpty(user.Name) || string.IsNullOrEmpty(user.Password))
            {
                throw new ArgumentException("Name and Password are required fields.");
            }
            user.Password = HashUtils.HashPassword(user.Password);
            user.UpdatedAt = DateTime.UtcNow;
            await _userRepository.AddUserAsync(user);
            return true;
        }

        public Task<LoginResponse> LoginEmailAsync(UserLogin user)
        {
            if (string.IsNullOrEmpty(user.Email) || string.IsNullOrEmpty(user.Password))
            {
                throw new ArgumentException("Email and Password are required for login.");
            }

            var existingUser = _userRepository.GetUserByEmailAsync(user.Email).Result;
            if (existingUser == null)
            {
                throw new InvalidOperationException("User not found.");
            }

            if (string.IsNullOrEmpty(existingUser.Password))
            {
                throw new InvalidOperationException("User password is not set.");
            }

            if (!HashUtils.VerifyPassword(user.Password, existingUser.Password))
            {
                throw new UnauthorizedAccessException("Invalid password.");
            }
            var token = _jwtUtils.GenerateToken(existingUser.Id.ToString());
            if (string.IsNullOrEmpty(token))
            {
                throw new InvalidOperationException("Failed to generate token.");
            }
            existingUser.Password = null; // Clear password before returning
            return Task.FromResult(new LoginResponse
            {
                User = existingUser,
                Token = token
            });
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