using Microsoft.AspNetCore.Mvc;
using SecureChat.src.service.UserService;
using SecureChat.src.db.schema;
using SecureChat.src.api.model;


namespace SecureChat.src.api.controller.v1.UserController
{
    [ApiController]
    [Route("api/v1/[controller]")]
    public class UserController(IUserService userService) : ControllerBase, IUserController
    {
        public IUserService _userService = userService;

        [HttpGet("/get/email/{email}")]
        public Task<IActionResult> GetUserByEmailAsync(string email)
        {
            throw new NotImplementedException();
        }

        [HttpGet("/get/{id}")]
        public Task<IActionResult> GetUserByIdAsync(int id)
        {
            throw new NotImplementedException();
        }

        [HttpGet("/get/phone/{phoneNumber}")]
        public Task<IActionResult> GetUserByPhoneNumberAsync(string phoneNumber)
        {
            throw new NotImplementedException();
        }

        [HttpPost("/login/email")]
        public Task<IActionResult> LoginEmailAsync([FromBody] UserLogin user)
        {
            if (user == null || string.IsNullOrEmpty(user.Email) || string.IsNullOrEmpty(user.Password))
            {
                return Task.FromResult<IActionResult>(new BadRequestObjectResult(new
                {
                    error = new
                    {
                        message = "Email and Password are required for login."
                    }
                }));
            }

            try
            {
                var loginResponse = _userService.LoginEmailAsync(user).Result;
                return Task.FromResult<IActionResult>(new OkObjectResult(loginResponse));
            }
            catch (Exception ex)
            {
                return Task.FromResult<IActionResult>(new BadRequestObjectResult(new
                {
                    error = new
                    {
                        message = ex.Message
                    }
                }));
            }
        }

        [HttpPost("/login/phone")]
        public Task<IActionResult> LoginPhoneNumberAsync([FromBody] UserLogin user)
        {
            throw new NotImplementedException();
        }

        [HttpPost("/register")]
        public async Task<IActionResult> RegisterAsync([FromBody] User user)
        {
            if (user == null)
            {
                return new BadRequestObjectResult("User cannot be null");
            }

            if (string.IsNullOrEmpty(user.Name) || string.IsNullOrEmpty(user.Password))
            {
                return new BadRequestObjectResult(new
                {
                    error = new
                    {
                        message = "Name and Password are required fields."
                    }
                });
            }

            try
            {
                var isRegister = await _userService.RegisterAsync(user);
                if (!isRegister)
                {
                    return new BadRequestObjectResult(new
                    {
                        error = new
                        {
                            message = "User registration failed."
                        }
                    });
                }
                return new OkObjectResult(new
                {
                    message = "User registered successfully.",
                    user = new
                    {
                        user.Id,
                        user.Name,
                        user.PhoneNumber,
                        user.Email,
                        user.CreatedAt,
                        user.UpdatedAt
                    }
                });
            }
            catch (Exception ex)
            {
                return new BadRequestObjectResult(new
                {
                    error = new
                    {
                        message = ex.Message
                    }
                });
            }
        }
    }
}