require 'socket'
socket = TCPSocket.open("localhost",9999)

1000.times do |i| 
  socket.write("#{i}\n")
end

1000.times {puts socket.gets}
 
