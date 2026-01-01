# README.md

run "jackpot_test" {
  command = plan
  assert {
    condition     = "jackpot\n" == output.jackpot_newline
    error_message = "gold.txt should contain jackpot not ${file(local.path)}"
  }
  assert {
    condition     = "jackpot" == output.jackpot
    error_message = "gold_no_newline.txt should contain jackpot not ${file(local.path_no_newline)}"
  }
  assert {
    condition     = "jackpot" == trimspace(output.jackpot_newline)
    error_message = "stripped gold.txt content should be \"jackpot\""
  }
}

run "minerals_test" {

  command = plan

  variables {
    available = {
      quartz   = false
      silicium = false
      obsidian = true
      redstone = true
    }
  }

  assert {
    condition     = alltrue([for mineral in output.minerals : lookup(var.available, mineral, false)])
    error_message = "only obsidian and redstone should be available"
  }

}

run "gems_test" {

  command = plan

  variables {
    available = {
      emerald = false
      ruby    = true
      diamond = true
    }
  }

  assert {
    condition     = alltrue([for gem, available in var.available : contains(output.gems, gem) == available])
    error_message = "only ruby and diamond should be available"
  }

}
