package util

//NonSolid 测试方块是否可以被穿过
func NonSolid(bs string) bool {
	return bs == "minecraft:air" || //空气
		bs == "minecraft:cave_air" ||

		bs == "minecraft:water" || //液体
		bs == "minecraft:lava" ||

		bs == "minecraft:oak_sapling" || //树苗
		bs == "minecraft:spruce_sapling" ||
		bs == "minecraft:birch_sapling" ||
		bs == "minecraft:jungle_sapling" ||
		bs == "minecraft:acacia_sapling" ||
		bs == "minecraft:dark_oak_sapling" ||

		bs == "minecraft:dandelion" || //花
		bs == "minecraft:poppy" ||
		bs == "minecraft:blue_orchid" ||
		bs == "minecraft:allium" ||
		bs == "minecraft:azure_bluet" ||
		bs == "minecraft:red_tulip" ||
		bs == "minecraft:orange_tulip" ||
		bs == "minecraft:white_tulip" ||
		bs == "minecraft:pink_tulip" ||
		bs == "minecraft:oxeye_daisy" ||
		bs == "minecraft:cornflower" || //1.14
		bs == "minecraft:lily_of_the_valley" || //1.14
		bs == "minecraft:wither_rose" || //1.14
		bs == "minecraft:sunflower" ||
		bs == "minecraft:lilac" ||
		bs == "minecraft:rose_bush" ||
		bs == "minecraft:peony" ||

		bs == "minecraft:torch" ||
		bs == "minecraft:wall_torch"
	//unfinished
}
