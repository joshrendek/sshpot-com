class Api::HttpRequestsController < ApplicationController
  skip_before_action :verify_authenticity_token
  def create
    body = Oj.load(request.body.read)
    return if body.nil?
    headers = body['headers'].map{|x| x.flatten.join(": ") }
    formdata =  body['form_data'].map{|x| { x[0] => x[1][0] } }
    HttpRequest.create(url: body['url'], response: body['response'].gsub("\u0000", ""),
                       hostname: body['hostname'], headers: headers,
                       formdata: formdata, method: body['method'], guid: body['guid'])
    render json: {status: :ok}, status: :ok
  end
end
