# The Color Chronicles

## Project Overview
*The Color Chronicles* is a 2D platformer game inspired by classic titles like *Super Mario*. Developed using the **Defold** engine with **Lua** scripting, the game offers dynamic levels, various types of enemies, collectibles, and engaging gameplay mechanics designed to provide a fun and challenging experience for players.

### Repository
[GitHub Repository](https://github.com/WojciechMierzwa/ProjektZaliczeniowy_Grafika_i_silniki_gier_komputerowych)

### Executable File
[Download the Executable (.exe)](https://anselbpl-my.sharepoint.com/:u:/g/personal/20858_student_ans-elblag_pl/EUnn9YpkS2ZPm7D9hP5PZTYB03B9g7HvOHGYOgAjoDp8RA)

## Features
- **Classic Platformer Gameplay**: Navigate through different levels, avoid obstacles, jump over gaps, and defeat enemies.
- **Boost Mode**: Temporarily become invincible after attacking an enemy.
- **Collectibles**: Gather coins to earn extra lives and purchase boosts in the in-game shop.
- **Boss Fights**: Face off against challenging bosses that require strategy and quick thinking.
- **Five Levels**: Each level features unique environments and color themes, with increasing difficulty.
- **Interactive Environment**: Players can activate objects, collect power-ups, and more.

## Controls
- **Arrow Left**: Move Left
- **Arrow Right**: Move Right
- **Space**: Jump
- **Q**: Return to the submenu
- **Left Click**: Interact with menus and buttons

### Menu Options
- **Main Menu**: Play, About, Quit, Reset (New Game)
- **Submenu**: Level Selection, Stats, Shop

## Game Rules
- **Boost Mode**: The player becomes invincible for 0.5 seconds after attacking an enemy. During this time, obstacles and base enemies won't harm the player.
- **Boss Fights**: The boss has two modes – Walk and Attack. The player can only attack the boss during its walk mode. The boss causes damage (2 lives) when in attack mode.
- **Level Completion**: Levels 1-4 are completed by reaching the flag at the end. Level 5 requires defeating the boss to complete the level.

## Enemies
- **Basic Enemies**: Moving monsters that can damage the player.
- **Traps**: Spikes and other hazards that the player must avoid to stay alive.
- **Boss**: The final enemy in Level 5, requiring strategic attacks to defeat.

## Game Strategies
- **Avoiding Attacks**: Predict enemy movements and avoid traps to preserve health.
- **Using Boosts**: Activate special abilities to deal with difficult enemies or to cross challenging sections.
- **Mastering Boss Fights**: Attack the boss during its walk phase while avoiding its powerful attack phase.

## Project Structure
- **assets/**: Contains all game resources (images, sounds, animations).
  - **character/**: Player-related assets.
  - **enemies/**: Various enemy types and their animations.
  - **icons/**: UI elements and menu icons.
  - **items/**: Collectible items like coins and boosts.
  - **platform_components/**: Platforms, obstacles, and interactive elements.
  - **sound/**: Sound effects and background music.

- **game/**: The core game logic and scripts.
  - **enemies/**: Different types of enemies, each with specific behaviors.
  - **game_controller/**: Manages game flow, level transitions, and game state.
  - **levels/**: Each level's design and logic, including assets and interactions.
  - **main/**: Main entry point to the game, linking all systems together.
  - **objects/**: Interactive objects like coins, boosts, and flags.
  - **player/**: Player-related logic and behaviors.

## Technologies Used
- **Defold Engine**: A powerful and flexible 2D game engine for creating cross-platform games.
- **Lua**: A lightweight scripting language used for game logic.
