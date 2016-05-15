Rails.application.routes.draw do
  resources :http_requests
  resources :logins

  resources :home do
    collection do
      get :logins_by_day, format: :json
      get :logins_by_country, format: :json
    end
  end

  namespace :api do
    resources :commands
    resources :logins
    resources :http_requests
  end
  root to: 'home#index'
end
