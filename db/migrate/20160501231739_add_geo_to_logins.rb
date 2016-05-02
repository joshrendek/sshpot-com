class AddGeoToLogins < ActiveRecord::Migration
  def change
    add_column :logins, :country_name, :string
    add_index :logins, :country_name
    add_column :logins, :country_code, :string
    add_index :logins, :country_code
  end
end
