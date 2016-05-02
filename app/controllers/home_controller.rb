class HomeController < ApplicationController
  def logins_by_day
    render json: Login.group_by_day(:created_at, range: 1.month.ago..Time.now, format: '%m/%d').count
  end

  def logins_by_country
    render json: Login.group(:country_name).where(created_at:  1.month.ago..Time.now).count
  end
end
