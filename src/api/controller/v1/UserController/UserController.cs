using Microsoft.AspNetCore.Mvc;
using SecureChat.src.service.UserService;
using SecureChat.src.db.schema;


namespace SecureChat.src.api.controller.v1.UserController
{
    [ApiController]
    [Route("api/v1/[controller]")]
    public class UserController(IUserService userService) : ControllerBase, IUserController
    {
        public IUserService _userService = userService;

        [HttpGet("/get/all")]
        public Task<IActionResult> GetAllUsersAsync()
        {
            throw new NotImplementedException();
        }

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
        public Task<IActionResult> LoginEmailAsync(string email, string password)
        {
            throw new NotImplementedException();
        }

        [HttpPost("/login/phone")]
        public Task<IActionResult> LoginPhoneNumberAsync(string phoneNumber, string password)
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
                var registeredUser = await _userService.RegisterAsync(user);
                return new OkObjectResult(registeredUser);
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