# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20160514022534) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "commands", force: :cascade do |t|
    t.text     "command"
    t.string   "guid"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  add_index "commands", ["command"], name: "index_commands_on_command", using: :btree
  add_index "commands", ["guid"], name: "index_commands_on_guid", using: :btree

  create_table "honeypots", force: :cascade do |t|
    t.inet     "ip"
    t.string   "guid"
    t.integer  "logins"
    t.string   "country_name"
    t.datetime "created_at",   null: false
    t.datetime "updated_at",   null: false
  end

  create_table "http_requests", force: :cascade do |t|
    t.string   "headers",    default: [],              array: true
    t.string   "url"
    t.string   "hostname"
    t.string   "formdata",   default: [],              array: true
    t.string   "method"
    t.string   "guid"
    t.datetime "created_at",              null: false
    t.datetime "updated_at",              null: false
    t.string   "response"
  end

  add_index "http_requests", ["guid"], name: "index_http_requests_on_guid", using: :btree
  add_index "http_requests", ["hostname"], name: "index_http_requests_on_hostname", using: :btree
  add_index "http_requests", ["method"], name: "index_http_requests_on_method", using: :btree
  add_index "http_requests", ["url"], name: "index_http_requests_on_url", using: :btree

  create_table "login_counts", force: :cascade do |t|
    t.inet     "ip"
    t.integer  "count",      default: 0
    t.datetime "created_at",             null: false
    t.datetime "updated_at",             null: false
  end

  add_index "login_counts", ["ip"], name: "index_login_counts_on_ip", using: :btree

  create_table "logins", force: :cascade do |t|
    t.inet     "remote_addr"
    t.integer  "remote_port"
    t.string   "username"
    t.string   "password"
    t.string   "guid"
    t.string   "version"
    t.string   "public_key"
    t.string   "key_type"
    t.string   "login_type"
    t.datetime "created_at",   null: false
    t.datetime "updated_at",   null: false
    t.string   "country_name"
    t.string   "country_code"
  end

  add_index "logins", ["country_code"], name: "index_logins_on_country_code", using: :btree
  add_index "logins", ["country_name"], name: "index_logins_on_country_name", using: :btree
  add_index "logins", ["guid"], name: "index_logins_on_guid", using: :btree
  add_index "logins", ["key_type"], name: "index_logins_on_key_type", using: :btree
  add_index "logins", ["login_type"], name: "index_logins_on_login_type", using: :btree
  add_index "logins", ["password"], name: "index_logins_on_password", using: :btree
  add_index "logins", ["remote_addr"], name: "index_logins_on_remote_addr", using: :btree
  add_index "logins", ["username"], name: "index_logins_on_username", using: :btree
  add_index "logins", ["version"], name: "index_logins_on_version", using: :btree

end
