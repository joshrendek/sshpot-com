class Api::HttpRequestsController < ApplicationController
  skip_before_action :verify_authenticity_token
  def create
    body = JSON.parse(request.body.read)
    headers = body['headers'].map{|x| x.flatten.join(": ") }
    formdata =  body['form_data'].map{|x| { x[0] => x[1][0] } }
    HttpRequest.create(url: body['url'], response: body['response'],
                       hostname: body['hostname'], headers: headers,
                       formdata: formdata, method: body['method'], guid: body['guid'])
    render json: {status: :ok}, status: :ok
  end
end