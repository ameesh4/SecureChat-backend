using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using SecureChat.src.service.UserService;
using SecureChat.src.db.schema;
using SecureChat.src.api.model;

namespace SecureChat.src.api.controller.v1.UserController
{
    public interface IUserController
    {
        Task<IActionResult> RegisterAsync(User user);
        Task<IActionResult> LoginEmailAsync([FromBody] UserLogin user);
        Task<IActionResult> LoginPhoneNumberAsync([FromBody] UserLogin user);
        Task<IActionResult> GetUserByIdAsync(int id);
        Task<IActionResult> GetUserByEmailAsync(string email);
        Task<IActionResult> GetUserByPhoneNumberAsync(string phoneNumber);
    }
}