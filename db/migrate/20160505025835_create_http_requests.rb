class CreateHttpRequests < ActiveRecord::Migration
  def change
    create_table :http_requests do |t|
      t.string :headers, array: true, default: []
      t.string :url
      t.string :hostname
      t.string :formdata, array: true, default: []
      t.string :method
      t.string :guid

      t.timestamps null: false
    end
    add_index :http_requests, :url
    add_index :http_requests, :hostname
    add_index :http_requests, :method
    add_index :http_requests, :guid
  end
end
