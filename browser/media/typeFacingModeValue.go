package media

type FacingModeValue string

func (e FacingModeValue) String() string {
	return string(e)
}

const (

	// KFacingModeUser
	//
	// English:
	//
	// The video source is facing toward the user; this includes, for example, the front-facing camera on a smartphone.
	//
	// Português:
	//
	// A fonte de vídeo está voltada para o usuário; isso inclui, por exemplo, a câmera frontal de um smartphone.
	KFacingModeUser FacingModeValue = "user"

	// KFacingModeEnvironment
	//
	// English:
	//
	// The video source is facing away from the user, thereby viewing their environment. This is the back camera on a
	// smartphone.
	//
	// Português:
	//
	// A fonte de vídeo está de costas para o usuário, visualizando assim seu ambiente. Esta é a câmera traseira em um
	// smartphone.
	KFacingModeEnvironment FacingModeValue = "environment"

	// KFacingModeLeft
	//
	// English:
	//
	// The video source is facing toward the user but to their left, such as a camera aimed toward the user but over their
	// left shoulder.
	//
	// Português:
	//
	// A fonte de vídeo está voltada para o usuário, mas à esquerda, como uma câmera voltada para o usuário, mas sobre o
	// ombro esquerdo.
	KFacingModeLeft FacingModeValue = "left"

	// KFacingModeRight
	//
	// English:
	//
	// The video source is facing toward the user but to their right, such as a camera aimed toward the user but over
	// their right shoulder.
	//
	// Português:
	//
	// A fonte de vídeo está voltada para o usuário, mas à direita, como uma câmera voltada para o usuário, mas sobre o
	// ombro direito.
	KFacingModeRight FacingModeValue = "right"
)
