class CreateHoneypots < ActiveRecord::Migration
  def change
    create_table :honeypots do |t|
      t.inet :ip
      t.string :guid
      t.integer :logins
      t.string :country_name

      t.timestamps null: false
    end
  end
end
