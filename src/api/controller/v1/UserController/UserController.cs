using Microsoft.AspNetCore.Mvc;
using hushline.src.service.UserService;
using hushline.src.api.controller.v1.UserController;
using hushline.src.db.schema;


namespace hushline.api.controller.v1.UserController
{
    [ApiController]
    [Route("api/v1/[controller]")]
    public class UserController : IUserController
    {
        public IUserService _userService;

        public UserController(IUserService userService)
        {
            _userService = userService;
        }

        public Task<IActionResult> GetAllUsersAsync()
        {
            throw new NotImplementedException();
        }

        public Task<IActionResult> GetUserByEmailAsync(string email)
        {
            throw new NotImplementedException();
        }

        public Task<IActionResult> GetUserByIdAsync(int id)
        {
            throw new NotImplementedException();
        }

        public Task<IActionResult> GetUserByPhoneNumberAsync(string phoneNumber)
        {
            throw new NotImplementedException();
        }

        public Task<IActionResult> LoginEmailAsync(string email, string password)
        {
            throw new NotImplementedException();
        }

        public Task<IActionResult> LoginPhoneNumberAsync(string phoneNumber, string password)
        {
            throw new NotImplementedException();
        }

        public Task<IActionResult> RegisterAsync(User user)
        {
            throw new NotImplementedException();
        }
    }
}