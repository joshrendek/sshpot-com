class HttpRequest < ActiveRecord::Base
  belongs_to :login, foreign_key: 'guid', primary_key: 'guid'
  def self.latest(n=25)
    order(id: :desc).limit(n)
  end

  def filtered_response
    ips = Honeypot.pluck(:ip).map(&:to_s)
    resp = response
    ips.each do |ip|
      resp.gsub!(ip, '*.*.*.*')
    end
    resp
  end
end
