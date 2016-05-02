class AddIndexCommands < ActiveRecord::Migration
  def change
    add_index :commands, :command
  end
end
