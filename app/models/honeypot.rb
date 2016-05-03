class Honeypot < ActiveRecord::Base
  before_create :set_guid
  before_create :set_country

  def self.top(n=10)
    order(logins: :desc).limit(n)
  end

  def set_country
    c = GeoIP.new("#{Rails.root}/GeoIP.dat").country(ip)
    self.country_name = c.country_name
  end

  def set_guid
    self.guid = SecureRandom.uuid
  end
end
