class Command < ActiveRecord::Base
  belongs_to :login, foreign_key: 'guid', primary_key: 'guid'

  def self.latest
    order(id: :desc).limit(25)
  end
end
