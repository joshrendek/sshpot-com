class LoginCount < ActiveRecord::Base

  def self.top
    order(count: :desc).limit(10)
  end

  def sessions
    Login.where(remote_addr: ip)
  end

  def commands
    Command.where(guid: Login.where(remote_addr: ip).pluck(:guid))
  end

  def self.inc(ip)
    a = find_or_create_by(ip: ip)
    a.increment(:count, 1)
    a.save
  end
end
