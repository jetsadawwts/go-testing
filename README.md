# Cmd
go test ./... = เปิดใช้งาน testing  ทั้งหมดตาม path บนโฟล์เดอร์
go test -v ./... = เปิดใช้งาน testing  ทั้งหมดตาม path บนโฟล์เดอร์ เเต่สามารถดู log ได้
got test . = เปิดใช้งาน testing ทั้งหมดภายใน Sub โฟล์เดอร์
got test -v . = เปิดใช้งาน testing ทั้งหมดภายใน Sub โฟล์เดอร์ เเต่สามารถดู log ได้



# Coding
t.log("") = เเสดง log
t.Error("") = เเสดง log Error
t.Errorf("") = เเสดง log Error หลายตัวเเปร
t.Run(string, func(t *testing.T)) = Sub test log

# other
Arrange, Act, Assent pattern
    

