locals {
  path            = "${path.module}/gold.txt"
  path_no_newline = "${path.module}/gold_no_newline.txt"
  gems            = fileset("${path.module}", "{ruby,diamond,emeral}.txt")
  minerals        = fileset("${path.module}", "{obsidian,redstone,quartz,silicium}.txt")
  confs           = fileset("${path.module}", "*.json")
}

output "jackpot_newline" {
  value = file(local.path)
}

output "jackpot" {
  value = file(local.path_no_newline)
}

output "minerals" {
  value = [for mineral in local.minerals : trimsuffix(mineral, ".txt")]
}

output "gems" {
  value = [for gem in local.gems : trimsuffix(gem, ".txt")]
}
