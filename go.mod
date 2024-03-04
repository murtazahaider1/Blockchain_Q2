module github.com/murtazahaider1/Blockchain_Q2

go 1.16

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency v0.1.0
)

require (
    github.com/additional/dependency v2.3.4 // indirect
)

replace github.com/some/dependency => github.com/forked/version v1.2.3
