using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using SecureChat.src.db.schema;

namespace SecureChat.src.db.repository.UserRepository
{
    public class UserRepository(AppDbContext context) : IUserRepository
    {
        private readonly AppDbContext _context = context;

        public async Task<User?> GetUserByIdAsync(int id)
        {
            var user = await _context.Users.FindAsync(id);
            return user ?? null;
        }

        public async Task<User?> GetUserByEmailAsync(string email)
        {
            var user = await _context.Users.FirstOrDefaultAsync(u => u.Email == email);
            User? user1 = user ?? null;
            return user1;
        }

        public async Task<User> GetUserByPhoneNumberAsync(string phoneNumber)
        {
            var user = await _context.Users.FirstOrDefaultAsync(u => u.PhoneNumber == phoneNumber);
            return user ?? throw new KeyNotFoundException($"User with phone number {phoneNumber} not found.");
        }

        public async Task<IEnumerable<User>> GetAllUsersAsync()
        {
            return await _context.Users.ToListAsync();
        }

        public async Task<User> AddUserAsync(User user)
        {
            await _context.Users.AddAsync(user);
            await _context.SaveChangesAsync();
            return user;
        }

        public async Task UpdateUserAsync(User user)
        {
            _context.Users.Update(user);
            await _context.SaveChangesAsync();
        }

        public async Task DeleteUserAsync(int id)
        {
            var user = await GetUserByIdAsync(id);
            if (user != null)
            {
                _context.Users.Remove(user);
                await _context.SaveChangesAsync();
            }
        }
    }
}