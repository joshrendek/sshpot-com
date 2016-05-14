class AddResponseToHttpRequests < ActiveRecord::Migration
  def change
    add_column :http_requests, :response, :string
  end
end
