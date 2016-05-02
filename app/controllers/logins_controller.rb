class LoginsController < ApplicationController
  def index
    @logins = Login.order(id: :desc)
    @logins = @logins.where(remote_addr: params[:ip]) if params[:ip].present?
    @logins = @logins.where(username: params[:username]) if params[:username].present?
    @logins = @logins.where(version: params[:client]) if params[:client].present?
    @logins = @logins.page(params[:page])
  end
end
