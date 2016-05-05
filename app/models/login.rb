class Login < ActiveRecord::Base
  paginates_per 25
  has_many :commands, foreign_key: 'guid', primary_key: 'guid'
  has_many :http_requests, foreign_key: 'guid', primary_key: 'guid'

  before_create :geolocate

  def geolocate
    c = GeoIP.new("#{Rails.root}/GeoIP.dat").country(remote_addr)
    self.country_code = c.country_code2
    self.country_name = c.country_name
  end

  def self.latest
    order('id desc').limit(25)
  end

  def self.latest_sessions
    joins(:commands).group('logins.id').having('count(commands.id) > 0').order(id: :desc)
  end
end
