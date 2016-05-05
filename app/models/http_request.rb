class HttpRequest < ActiveRecord::Base
  belongs_to :login, foreign_key: 'guid', primary_key: 'guid'
  def self.latest(n=25)
    order(id: :desc).limit(n)
  end
end
