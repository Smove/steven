package main

// Valid blocks.
var (
	BlockAir                        = initSimpleConfig("air", simpleConfig{NotCullAgainst: true})
	BlockStone                      = initStone("stone")
	BlockGrass                      = initGrass()
	BlockDirt                       = initSimple("dirt")
	BlockCobblestone                = initSimple("cobblestone")
	BlockPlanks                     = initSimple("planks")
	BlockSapling                    = initSimpleConfig("sapling", simpleConfig{NotCullAgainst: true})
	BlockBedrock                    = initSimple("bedrock")
	BlockFlowingWater               = initLiquid("flowing_water", false)
	BlockWater                      = initLiquid("water", false)
	BlockFlowingLava                = initLiquid("flowing_lava", true)
	BlockLava                       = initLiquid("lava", true)
	BlockSand                       = initSimple("sand")
	BlockGravel                     = initSimple("gravel")
	BlockGoldOre                    = initSimple("gold_ore")
	BlockIronOre                    = initSimple("iron_ore")
	BlockCoalOre                    = initSimple("coal_ore")
	BlockLog                        = initLog("log")
	BlockLeaves                     = initLeaves("leaves")
	BlockSponge                     = initSimple("sponge")
	BlockGlass                      = initSimpleConfig("glass", simpleConfig{NotCullAgainst: true})
	BlockLapisOre                   = initSimple("lapis_ore")
	BlockLapisBlock                 = initSimple("lapis_block")
	BlockDispenser                  = initSimple("dispenser")
	BlockSandstone                  = initSimple("sandstone")
	BlockNote                       = initSimple("noteblock")
	BlockBed                        = initSimple("bed")
	BlockGoldenRail                 = initSimple("golden_rail")
	BlockDetectorRail               = initSimple("detector_rail")
	BlockStickyPiston               = initSimple("sticky_piston")
	BlockWeb                        = initSimple("web")
	BlockTallGrass                  = initSimple("tallgrass")
	BlockDeadBush                   = initSimple("deadbush")
	BlockPiston                     = initSimple("piston")
	BlockPistonHead                 = initSimple("piston_head")
	BlockWool                       = initSimple("wool")
	BlockPistonExtension            = initSimple("piston_extension")
	BlockYellowFlower               = initSimple("yellow_follow")
	BlockRedFlower                  = initSimple("red_flower")
	BlockBrownMushroom              = initSimple("brown_mushroom")
	BlockRedMushrrom                = initSimple("red_mushroom")
	BlockGoldBlock                  = initSimple("gold_block")
	BlockIronBlock                  = initSimple("iron_block")
	BlockDoubleStoneSlab            = initSimple("double_stone_slab")
	BlockStoneSlab                  = initSimple("stone_slab")
	BlockBrickBlock                 = initSimple("brick_block")
	BlockTNT                        = initSimple("tnt")
	BlockBookShelf                  = initSimple("blockshelf")
	BlockMossyCobblestone           = initSimple("mossy_cobblestone")
	BlockObsidian                   = initSimple("obsidian")
	BlockTorch                      = initSimple("torch")
	BlockFire                       = initSimple("fire")
	BlockMobSpawner                 = initSimple("mob_spawner")
	BlockOakStairs                  = initSimple("oak_stairs")
	BlockChest                      = initSimple("chest")
	BlockRedstoneWire               = initSimple("redstone_wire")
	BlockDiamondOre                 = initSimple("diamond_ore")
	BlockDiamondBlock               = initSimple("diamond_block")
	BlockCraftingTable              = initSimple("crafting_table")
	BlockWheat                      = initSimple("wheat")
	BlockFarmland                   = initSimple("farmland")
	BlockFurnace                    = initSimple("furnace")
	BlockFurnaceLit                 = initSimple("lit_furnace")
	BlockStandingSign               = initSimple("standing_sign")
	BlockWoodenDoor                 = initSimple("wooden_door")
	BlockLadder                     = initSimple("ladder")
	BlockRail                       = initSimple("rail")
	BlockStoneStairs                = initSimple("stone_stairs")
	BlockWallSign                   = initSimple("wall_sign")
	BlockLever                      = initSimple("lever")
	BlockStonePressurePlate         = initSimple("stone_pressure_plate")
	BlockIronDoor                   = initSimple("iron_door")
	BlockWoodenPressurePlate        = initSimple("wooden_pressure_plate")
	BlockRedstoneOre                = initSimple("redstone_ore")
	BlockRedstoneOreLit             = initSimple("lit_redstone_ore")
	BlockRedstoneTorchUnlit         = initSimple("unlit_redstone_torch")
	BlockRedstoneTorch              = initSimple("redstone_torch")
	BlockStoneButton                = initSimple("stone_button")
	BlockSnowLayer                  = initSimple("snow_layer")
	BlockIce                        = initSimple("ice")
	BlockSnow                       = initSimple("snow")
	BlockCactus                     = initSimpleConfig("cactus", simpleConfig{NotCullAgainst: true})
	BlockClay                       = initSimple("clay")
	BlockReeds                      = initSimpleConfig("reeds", simpleConfig{NotCullAgainst: true})
	BlockJukebox                    = initSimple("jukebox")
	BlockFence                      = initSimple("fence")
	BlockPumpkin                    = initSimple("pumpkin")
	BlockNetherrack                 = initSimple("netherrack")
	BlockSoulSand                   = initSimple("soul_sand")
	BlockGlowstone                  = initSimple("glowstone")
	BlockPortal                     = initSimple("portal")
	BlockPumpkinLit                 = initSimple("lit_pumpkin")
	BlockCake                       = initSimple("cake")
	BlockRepeaterUnpowered          = initSimple("unpowered_repeater")
	BlockRepeaterPowered            = initSimple("powered_repeater")
	BlockStainedGlass               = initSimple("stained_glass")
	BlockTrapDoor                   = initSimple("trapdoor")
	BlockMonsterEgg                 = initSimple("monster_egg")
	BlockStoneBrick                 = initSimple("stonebrick")
	BlockBrownMushroomBlock         = initSimple("brown_mushroom_block")
	BlockRedMushroomBlock           = initSimple("red_mushroom_block")
	BlockIronBars                   = initSimple("iron_bars")
	BlockGlassPane                  = initSimple("glass_pane")
	BlockMelonBlock                 = initSimple("melon_block")
	BlockPumpkinStem                = initSimple("pumpkin_stem")
	BlockMelonStem                  = initSimple("melon_stem")
	BlockVine                       = initSimple("vine")
	BlockFenceGate                  = initSimple("fence_gate")
	BlockBrickStairs                = initSimple("brick_stairs")
	BlockStoneBrickStairs           = initSimple("stone_brick_stairs")
	BlockMycelium                   = initSimple("mycelium")
	BlockWaterlily                  = initSimple("waterlily")
	BlockNetherBrick                = initSimple("nether_brick")
	BlockNetherBrickFence           = initSimple("nether_brick_fence")
	BlockNetherBrickStairs          = initSimple("nether_brick_stairs")
	BlockNetherWart                 = initSimple("nether_wart")
	BlockEnchantingTable            = initSimple("enchanting_table")
	BlockBrewingStand               = initSimple("brewing_stand")
	BlockCauldron                   = initSimple("cauldron")
	BlockEndPortal                  = initSimple("end_portal")
	BlockEndPortalFrame             = initSimple("end_portal_frame")
	BlockDragonEgg                  = initSimple("dragon_egg")
	BlockRedstoneLamp               = initSimple("redstone_lamp")
	BlockRedstoneLampLit            = initSimple("lit_redstone_lamp")
	BlockDoubleWoodenSlab           = initSimple("double_wooden_slab")
	BlockWoodenSlab                 = initSimple("wooden_slab")
	BlockCocoa                      = initSimple("cocoa")
	BlockSandstoneStairs            = initSimple("sandstone_stairs")
	BlockEmeraldOre                 = initSimple("emerald_ore")
	BlockEnderChest                 = initSimple("ender_Chest")
	BlockTripwireHook               = initSimple("tripwire_hook")
	BlockTripwire                   = initSimple("tripwire")
	BlockEmeraldBlock               = initSimple("emerald_block")
	BlockSpruceStairs               = initSimple("spruce_stairs")
	BlockBirchStairs                = initSimple("birch_stairs")
	BlockJungleStairs               = initSimple("jungle_stairs")
	BlockCommandBlock               = initSimple("command_block")
	BlockBeacon                     = initSimple("beacon")
	BlockCobblestoneWall            = initSimple("cobblestone_wall")
	BlockFlowerPot                  = initSimple("flower_pot")
	BlockCarrots                    = initSimple("carrots")
	BlockPotatoes                   = initSimple("potatoes")
	BlockWoodenButton               = initSimple("wooden_button")
	BlockSkull                      = initSimple("skull")
	BlockAnvil                      = initSimple("anvil")
	BlockTrappedChest               = initSimple("trapped_chest")
	BlockLightWeightedPressurePlate = initSimple("light_weighted_pressure_plate")
	BlockHeavyWeightedPressurePlate = initSimple("heavy_weighted_pressure_plate")
	BlockComparatorUnpowered        = initSimple("unpowered_comparator")
	BlockComparatorPowered          = initSimple("powered_comparator")
	BlockDaylightDetector           = initSimple("daylight_detector")
	BlockRedstoneBlock              = initSimple("redstone_block")
	BlockQuartzOre                  = initSimple("quartz_ore")
	BlockHopper                     = initSimple("hopper")
	BlockQuartzBlock                = initSimple("quartz_block")
	BlockQuartzStairs               = initSimple("quartz_stairs")
	BlockActivatorRail              = initSimple("activator_rail")
	BlockDropper                    = initSimple("dropper")
	BlockStainedHardenedClay        = initSimple("stained_hardened_clay")
	BlockStainedGlassPane           = initSimple("stained_glass_pane")
	BlockLeaves2                    = initSimple("leaves2")
	BlockLog2                       = initSimple("log2")
	BlockAcaciaStairs               = initSimple("acacia_stairs")
	BlockDarkOakStairs              = initSimple("dark_oak_stairs")
	BlockSlime                      = initSimple("slime")
	BlockBarrier                    = initSimple("barrier")
	BlockIronTrapDoor               = initSimple("iron_trapdoor")
	BlockPrismarine                 = initSimple("prismarine")
	BlockSeaLantern                 = initSimple("sea_lantern")
	BlockHayBlock                   = initSimple("hay_block")
	BlockCarpet                     = initSimple("carpet")
	BlockHardenedClay               = initSimple("hardened_clay")
	BlockCoalBlock                  = initSimple("coal_block")
	BlockPackedIce                  = initSimple("packed_ice")
	BlockDoublePlant                = initSimple("double_plant")
	BlockStandingBanner             = initSimple("standing_banner")
	BlockWallBanner                 = initSimple("wall_banner")
	BlockDaylightDetectorInverted   = initSimple("daylight_detector_inverted")
	BlockRedSandstone               = initSimple("red_sandstone")
	BlockRedSandstoneStairs         = initSimple("red_sandstone_stairs")
	BlockDoubleStoneSlab2           = initSimple("double_stone_slab2")
	BlockStoneSlab2                 = initSimple("stone_slab2")
	BlockSpruceFenceGate            = initSimple("spruce_fence_gate")
	BlockBirchFenceGate             = initSimple("birch_fence_gate")
	BlockJungleFenceGate            = initSimple("jungle_fence_gate")
	BlockDarkOakFenceGate           = initSimple("dark_oak_fence_gate")
	BlockAcaciaFenceGate            = initSimple("acacia_fence_gate")
	BlockSpruceFence                = initSimple("spruce_fence")
	BlockBirchFence                 = initSimple("birch_fence")
	BlockJungleFence                = initSimple("jungle_fence")
	BlockDarkOakFence               = initSimple("dark_oak_fence")
	BlockAcaciaFence                = initSimple("acacia_fence")
	BlockSpruceDoor                 = initSimple("spruce_door")
	BlockBirchDoor                  = initSimple("birch_door")
	BlockJungleDoor                 = initSimple("jungle_door")
	BlockDarkOakDoor                = initSimple("dark_oak_door")
	BlockAcaciaDoor                 = initSimple("acacia_door")
)
