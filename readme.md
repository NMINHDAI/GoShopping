# Backend

GOLANG

### List of model

**A. User System**

- ID: ObjectID
- Role (user / admin / vendor): enum
- Email: string email
- Password: string password
- maybe Extra attribute : CreateAt, modifyAt,...

**B. Order System**

- ID: ObjectID
- List of tuple (IDProduct, Quantity)
- IDUser
- PaymentInfo (default: Null)

Note: nếu paymentInfo = Null ⇒ đang ở giỏ hàng, else nó đã được thanh toán và là đơn hàng.

**C. Product**

- ID: ObjectID
- Name: string
- Image (maybe): base64 or maybe string
- Discription: string
- Price: float

**D. Shop** 

- List of tuple (IDProduct, Quantity)

### **List of API (Route)**

**A. User System**

POST /user/login {email, password} return {UserID, jsonwebtoken}

POST /user/signup {email, password}

POST /user/forgot-password {email}

**B. Product System**

Khi đăng nhập và đăng kí xong, 1 web landing page sẽ call API, map product, và hiển thị lên UI:

GET / 

Người dùng có thể view chi tiết từng Product

GET /products/:IDProduct

Search bar given pattern matching ⇒ lấy ra tất cả các product theo điều kiện nào đó

GET /products/:pattern

Giỏ hàng để cập nhật realtime: ⇒ dùng useState của React, tuy nhiên ví dụ người A mua xong thì bên dashboard của người B phải ngay lập tức update lại quantity mà không cần reload

Em dự tính dùng socket và add tất cả quantity vô 1 kênh.

Khi người dùng nhấn mua hàng trong giỏ hàng:

POST /order/ 

{

UserID, 

listTuple(IDProduct, Quantity),

paymentInfo,

}