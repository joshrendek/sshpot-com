class Api::LoginsController < ApplicationController
  skip_before_action :verify_authenticity_token
  def create
    body = JSON.parse(request.body.read)
    Login.create(body)
    LoginCount.inc(body['remote_addr'])
    render json: {status: :ok}, status: :ok
  end
end
