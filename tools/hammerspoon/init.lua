-- local vpnIsConnected = function(vpnName)
-- 	local vpnStatus = hs.execute(string.format("networksetup -showpppoestatus \"%s\"", vpnName))
-- 	return vpnStatus == "connected\n"
-- end
-- local vpnNmae = "Geekbang"

-- -- 切换 Geekbang DNS
-- -- 先执行 sudo mkdir /etc/resolver && sudo chmod 777 /etc/resolver

-- local geekDnsBar = hs.menubar.new()
-- local geekDnsFile = "/etc/resolver/geekbang.org"
-- local geekCommonDnsHash = {
-- 	{"223.5.5.5","G"}
-- }
-- local geekVpnDnsHash = {
-- 	{"172.72.1.1","GP"},
-- 	{"172.72.2.2","GT"},
-- 	{"172.72.3.3","GS"}
-- }
-- local geekVpnGateIP = "192.168.42.1"
-- local geekDnsHash = geekCommonDnsHash
-- local geekDnsI = 1
-- hs.execute(string.format("echo \"nameserver %s\nsearch_order 1\ntimeout 5\" > %s", geekDnsHash[geekDnsI][1], geekDnsFile))

-- function geekDnsInit(flags)
-- 	if vpnIsConnected(vpnNmae) or (flags and flags & 1 > 0) then
-- 		-- VPN tunnel is up
-- 		geekDnsHash = geekVpnDnsHash
-- 		geekDnsI = 3
-- 	else
-- 		-- VPN tunnel is down
-- 		geekDnsHash = geekCommonDnsHash
-- 	end
-- 	--hs.alert.show(flags)
-- 	geekDnsChange()
-- end

-- function geekDnsChange()
-- 	geekDnsI = geekDnsI%#geekDnsHash+1
-- 	hs.execute(string.format("sed -i \"\" \"s/nameserver *.*.*.*/nameserver %s/g\" %s", geekDnsHash[geekDnsI][1], geekDnsFile))
-- 	geekDnsBar:setTitle(geekDnsHash[geekDnsI][2])
-- end

-- --
-- geekDnsInit()
-- hs.network.reachability.forAddress(geekVpnGateIP):setCallback(function(self, flags)
-- 	geekDnsInit(flags)
-- end):start()

-- hs.hotkey.bind({"cmd"}, "G", geekDnsChange)
-- geekDnsBar:setClickCallback(geekDnsChange)

-- 工具箱
function utils()
	local pastContent = hs.pasteboard.getContents()
	local toastContent = "^_^"
	local pastContentInt = tonumber(pastContent)
	
	-- 时间戳转换
	if string.len(pastContent) == 10 and pastContentInt ~= 0 then
        toastContent = os.date("%Y-%m-%d %X", pastContentInt)
        hs.messages.SMS("15232326681", toastContent)
	end
	-- 毫秒时间戳转换
	if string.len(pastContent) == 13 and pastContentInt ~= 0 then
		toastContent = os.date("%Y-%m-%d %X", math.floor(pastContentInt/1000))
		toastContent = toastContent.." "..tostring(pastContentInt%1000)
	end
	-- 用户信息
	--curl https://admin.geekbang.org/bill/user/info -d '{"cellphone":"13467676577"}' --cookie dx_token=d4451c6967de2b74f4083bf0760dd8f2
	hs.alert.show(toastContent)
end
hs.hotkey.bind({"cmd"}, "U", utils)


-- -- VPN
-- function vpnDisConnect(vpnName)
-- 	if vpnIsConnected(vpnName) then
-- 		hs.execute(string.format("networksetup -disconnectpppoeservice \"%s\"", vpnName))
-- 	end
-- end

-- function vpnConnect(vpnName)
-- 	if vpnIsConnected(vpnName) == false then
-- 		hs.execute(string.format("networksetup -connectpppoeservice \"%s\"", vpnName))
-- 	end
-- end

-- -- 应用
-- function appKill(name)
-- 	local app = hs.application.get(name)
-- 	if app and app:isRunning() then
-- 		app:kill()
-- 	end
-- end

-- function appLaunch(name)
-- 	hs.application.launchOrFocus(name)
-- end

-- wifi变更
-- local workWifi = 'Geekbang'
-- local homeWifi = 'yantu_5G'
-- hs.wifi.watcher.new(function()
--     local currentWifi = hs.wifi.currentNetwork()
-- 	
-- 	if not currentWifi or currentWifi == homeWifi then
-- 		appKill("V2rayU")
-- 	else
-- 		os.execute("sleep 2")
-- 		appLaunch("V2rayU")
-- 	end
-- 	
-- 	if currentWifi == workWifi then
-- 		os.execute("sleep 2")
-- 		vpnConnect("Geekbang")
-- 	end
-- 	
-- end):start()


