class HttpRequestsController < ApplicationController
  def index
    @requests = HttpRequest.order(id: :desc)
    @requests = @requests.page(params[:page])
  end

  def show
    @request = HttpRequest.find(params[:id])
    @login = @request.try(:login)
  end
end
