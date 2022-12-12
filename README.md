# Cmd
go test ./... = เปิดใช้งาน testing  ทั้งหมดตาม path บนโฟล์เดอร์
go test -v ./... = เปิดใช้งาน testing  ทั้งหมดตาม path บนโฟล์เดอร์ เเต่สามารถดู log ได้
go test . = เปิดใช้งาน testing ทั้งหมดภายใน Sub โฟล์เดอร์
go test -v . = เปิดใช้งาน testing ทั้งหมดภายใน Sub โฟล์เดอร์ เเต่สามารถดู log ได้
go test -v = เปิดใช้งาน testing โฟล์เดอร์ เเต่สามารถดู log ได้
go test -cover = Run test with converage เช็ดว่าเราเทสได้ครบทุกเงือนไขเเล้วหรือยัง
go test -v -cover = Run test with converage เช็ดว่าเราเทสได้ครบทุกเงือนไขเเล้วหรือยัง เเต่สามารถดู log ได้
go test -cover -coverprofile=c.out = Generating an c.out report
go tool cover -html=c.out -o coverage.html = Generating c.out report to HTML converage report
open [filename] = เปิดไฟล์
go fmt = จัด Format
go vet = คำสั่งที่ช่วยหา คำสั่งที่คิดว่าจะผิด หรือมีข้อผิดพลาด
go lint = ช่วยจัดการว่าเรานั้นเขียนได้ดีเเค่ไหน ค่อยจัดการหรือเเนะนำในการเขียน
staticcheck ./... = รัน Staticcheck
go doc = godoc -http=:6060
go build = go build -tags=name
go test 
    - go test -tags=integration
    - go test -v -tags integration
    - go test -v -tags db
    - go test -v -tags integration, db


# Coding
t.log("") = เเสดง log
t.Error("") = เเสดง log Error
t.Errorf("") = เเสดง log Error หลายตัวเเปร
t.Run(string, func(t *testing.T)) = Sub test log
... = space out คือการกระจายออกไป Ex got := sum([]int{2, 3, 3, -1}...) -> 2, 3, 3, -1
# other
Arrange, Act, Assent pattern
กรณีอยากจะตรวจสอบความถูกต้องของ integer ต้องใช้  uint
Setup / Teardown = เทสกรณีที่มีการต่อ Database หรือ Api


# Command Code
F2 = เปลี่ยนชื่อตัวเเปร

# install
gostaticcheck = go install -v honnef.co/go/tools/cmd/staticcheck@latest
godoc = go install -v golang.org/x/tools/cmd/godoc@latest

# link url
go test = https://pkg.go.dev/testing

    

