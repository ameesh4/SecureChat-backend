using Microsoft.AspNetCore.Mvc;
using SecureChat.api.model;

namespace SecureChat.api.controller.v1.UserController
{
    public interface IUserController
    {
        Task<IActionResult> RegisterAsync(Register user);
        Task<IActionResult> LoginEmailAsync([FromBody] UserLoginEmail user);
        Task<IActionResult> LoginPhoneNumberAsync([FromBody] UserLoginPhone user);
    }
}