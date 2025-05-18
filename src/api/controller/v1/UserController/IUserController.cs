using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using hushline.src.service.UserService;
using hushline.src.db.schema;

namespace hushline.src.api.controller.v1.UserController
{
    public interface IUserController
    {
        Task<IActionResult> RegisterAsync(User user);
        Task<IActionResult> LoginEmailAsync(string email, string password);
        Task<IActionResult> LoginPhoneNumberAsync(string phoneNumber, string password);
        Task<IActionResult> GetUserByIdAsync(int id);
        Task<IActionResult> GetUserByEmailAsync(string email);
        Task<IActionResult> GetUserByPhoneNumberAsync(string phoneNumber);
        Task<IActionResult> GetAllUsersAsync();
    }
}