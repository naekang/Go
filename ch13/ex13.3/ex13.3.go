package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

// 구조체에서 다른 구조체를 필드로 포함할 때 필드명을 생략하면 .을 한번만 찍어 접근 가능
func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)

	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
