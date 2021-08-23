Gem::Specification.new do |s|
  s.name        = 'outsert'
  s.version     = '0.0.0'
  s.executables << 'outsert'
  s.summary     = "outsert from (select * from users)"
  s.description = "Generate Insert Statement from select statement result csv"
  s.authors     = ["SUGI Kohei"]
  s.email       = 'koheisg@hey.com'
  s.files       = ["lib/outsert.rb", "bin/outsert"]
  s.homepage    = 'https://rubygems.org/gems/outsert'
  s.license     = 'MIT'
end
