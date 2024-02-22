# campfinderをクリーンアーキテクチャで実装
[campfinderのリポジトリ(2024/2/11)](https://github.com/tusmasoma/campfinder)をクリーンアーキテクチャに沿って実装しました。

クリーンアーキテクチャという概念をどう捉えるかについては、大きく分けて二つの意見があります。一つは、以下の図に登場する用語や概念を用いて、4層レイヤーからなるAPIアーキテクチャを指す場合です。もう一つは、クリーンアーキテクチャで重要とされる「依存性のルール」、「関心の分離」、「依存性逆転の原則(DIP)」などに着目し、図にはこだわらない方がいいという意見です。実際に、クリーンアーキテクチャの提案者のRobert C. Martin氏は、「図は概要であり、クリーンアーキテクチャは4層以上になりうる」と述べています。

今回は、前者の定義を採用します。すなわち、クリーンアーキテクチャとは以下の図に登場する用語や概念を用いて、4層レイヤーからなるAPIアーキテクチャを指すものとします。

![クリーンアーキテクチャ](https://github.com/tusmasoma/clean-architecture-campfinder/assets/104899572/ce8caa07-36ff-4d97-b201-5d559169eabc)

## 今回の実装の注意点
- RepositoryはUsecase層が持つことにします
  - RepositoryをEntity層が持つこともできます

- 今回は、上記の図をできるだけ正確に実装する為、処理のフローも上記の図の右下に従います。つまり、InputPort、OutputPortを実装し、それを経由して出力をUsecase層のユースケースで実行することにします([#v1.0.0](https://github.com/tusmasoma/clean-architecture-campfinder/releases/tag/v.1.0.0))

  例)
  
  ```go
  package interactor
  
  type User struct {
      OutputPort port.UserOutputPort
      UserRepo   port.UserRepository
  }

  func (u *User) GetUserByID(ctx context.Context, userID string) {
      user, err := u.UserRepo.GetUserByID(ctx, userID)

      if err != nil {
          u.OutputPort.RenderError(err)
	  return
      }
  
      u.OutputPort.Render(user) // Usecase層でOutputPortを経由してpresenter実行
  }
  ```

  ```go
  package controller
  
  type User struct {
	    OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	    // -> presenter.NewUserOutputPort
	    InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	    // -> interactor.NewUserInputPort
	    RepoFactory func(c *sql.DB) port.UserRepository
	    // -> gateway.NewUserRepository
	    Conn *sql.DB
  }

  func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
      ctx := r.Context()
      userID := strings.TrimPrefix(r.URL.Path, "/user/")
      outputPort := u.OutputFactory(w)
      repository := u.RepoFactory(u.Conn)
      inputPort := u.InputFactory(outputPort, repository)
      inputPort.GetUserByID(ctx, userID) // controllerでInputPort経由してUsecase層のinteractorを実行
  }
  ```

- [#v1.1.0](https://github.com/tusmasoma/clean-architecture-campfinder/compare/v.1.0.0...v1.1.0)では、コントロールラーにてoutputPortの出力を実行するように変更しました。そちらのほうが、本来のユースケース層とコントローラの役目と合っているからです。
