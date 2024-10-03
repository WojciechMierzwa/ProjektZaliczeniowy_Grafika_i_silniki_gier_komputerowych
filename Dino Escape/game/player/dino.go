components {
  id: "player"
  component: "/game/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/dino_sprite/DinoSprites.tilesource\"\n"
  "}\n"
  ""
}
