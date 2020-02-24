module github.com/stashapp/stash

require (
	github.com/99designs/gqlgen v0.9.0
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/antchfx/htmlquery v1.2.0
	github.com/antchfx/xpath v1.1.2 // indirect
	github.com/bmatcuk/doublestar v1.1.5
	github.com/disintegration/imaging v1.6.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/gobuffalo/packr/v2 v2.0.2
	github.com/golang-migrate/migrate/v4 v4.3.1
	github.com/gorilla/websocket v1.4.0
	github.com/h2non/filetype v1.0.8
	github.com/jmoiron/sqlx v1.2.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/rs/cors v1.6.0
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/vektah/gqlparser v1.1.2
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/image v0.0.0-20190118043309-183bebdce1b2 // indirect
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
	gopkg.in/yaml.v2 v2.2.2
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

go 1.13
