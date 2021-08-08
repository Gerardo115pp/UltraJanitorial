module Txy_server

replace (
	github.com/Gerardo115pp/PatriotLib/PatriotEcho => /home/el_maligno/SoftwareProjects/Development_resources/go_modules/patriotslibs.com/patriot-utils/
	github.com/Gerardo115pp/PatriotLib/PatriotFs => /home/el_maligno/SoftwareProjects/Development_resources/go_modules/patriotslibs.com/PatriotsFs/
	github.com/Gerardo115pp/PatriotLib/PatriotRouter => /home/el_maligno/SoftwareProjects/Development_resources/go_modules/patriotslibs.com/Router/
)

go 1.16

require (
	github.com/Gerardo115pp/PatriotLib/PatriotEcho v0.0.0-00010101000000-000000000000
	github.com/Gerardo115pp/PatriotLib/PatriotFs v0.0.0-00010101000000-000000000000
	github.com/Gerardo115pp/PatriotLib/PatriotRouter v0.0.0-00010101000000-000000000000
)
