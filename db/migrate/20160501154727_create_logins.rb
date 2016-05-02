class CreateLogins < ActiveRecord::Migration
  def change
    create_table :logins do |t|
      t.inet :remote_addr
      t.integer :remote_port
      t.string :username
      t.string :password
      t.string :guid
      t.string :version
      t.string :public_key
      t.string :key_type
      t.string :login_type

      t.timestamps null: false
    end
    add_index :logins, :remote_addr
    add_index :logins, :username
    add_index :logins, :password
    add_index :logins, :guid
    add_index :logins, :version
    add_index :logins, :key_type
    add_index :logins, :login_type
  end
end
