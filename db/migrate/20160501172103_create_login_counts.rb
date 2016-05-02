class CreateLoginCounts < ActiveRecord::Migration
  def change
    create_table :login_counts do |t|
      t.inet :ip
      t.integer :count, default: 0

      t.timestamps null: false
    end
    add_index :login_counts, :ip
  end
end
