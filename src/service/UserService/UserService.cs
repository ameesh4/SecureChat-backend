using SecureChat.api.model;
using SecureChat.db.repository.UserRepository;
using SecureChat.db.schema;
using SecureChat.service.JwtService;
using SecureChat.utils;

namespace SecureChat.service.UserService
{
    public class UserService(IUserRepository userRepository, IJwtService jwtService) : IUserService
    {
        private readonly IUserRepository _userRepository = userRepository;
        private readonly IJwtService _jwtUtils = jwtService;

        async Task<bool> IUserService.RegisterAsync(Register user)
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
            var userDb = new User
            {
                Name = user.Name,
                Email = user.Email,
                PhoneNumber = user.PhoneNumber,
                Password = HashUtils.HashPassword(user.Password),
                CreatedAt = DateTime.UtcNow,
                UpdatedAt = DateTime.UtcNow
            };
            await _userRepository.AddUserAsync(userDb);
            return true;
        }

        public Task<LoginResponse> LoginEmailAsync(UserLoginEmail user)
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
            var refreshToken = _jwtUtils.GenerateRefreshToken(existingUser.Id.ToString());
            var token = _jwtUtils.GenerateToken(existingUser.Id.ToString());
            existingUser.RefreshToken = refreshToken;
            if (string.IsNullOrEmpty(token))
            {
                throw new InvalidOperationException("Failed to generate token.");
            }
            _userRepository.UpdateUserAsync(existingUser).Wait(); 
            existingUser.Password = null; // Clear password before returning
            return Task.FromResult(new LoginResponse
            {
                User = existingUser,
                Token = token
            });
        }

        public Task<LoginResponse> LoginPhoneNumberAsync(UserLoginPhone user)
        {
            if (string.IsNullOrEmpty(user.PhoneNumber) || string.IsNullOrEmpty(user.Password))
            {
                throw new ArgumentException("Phone number and Password are required for login.");
            }

            var existingUser = _userRepository.GetUserByPhoneNumberAsync(user.PhoneNumber).Result;
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
            existingUser.RefreshToken = token;
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