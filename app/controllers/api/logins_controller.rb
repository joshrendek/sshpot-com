class Api::LoginsController < ApplicationController
  skip_before_action :verify_authenticity_token
  def create
    body = Oj.load(request.body.read)
    ip = request.env['HTTP_X_FORWARDED_FOR'] || request.remote_ip
    Rails.logger.info "******************"
    Rails.logger.info ip
    Rails.logger.info "******************"
    honeypot = Honeypot.find_or_create_by(ip: ip)
    honeypot.increment!(:logins, 1)
    Login.create(body)
    LoginCount.inc(body['remote_addr'])
    render json: {status: :ok}, status: :ok
  end
end
