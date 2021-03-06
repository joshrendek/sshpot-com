class Api::CommandsController < ApplicationController
  skip_before_action :verify_authenticity_token
  def create
    body = Oj.load(request.body.read)
    Command.create(body)
    render json: {status: :ok}, status: :ok
  end
end
