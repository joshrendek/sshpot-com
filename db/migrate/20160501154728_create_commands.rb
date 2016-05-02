class CreateCommands < ActiveRecord::Migration
  def change
    create_table :commands do |t|
      t.text :command
      t.string :guid

      t.timestamps null: false
    end
    add_index :commands, :guid
  end
end
