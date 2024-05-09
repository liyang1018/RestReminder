package text

import (
	"fyne/config"
	"math/rand"
	"time"
)

var remindText = make(map[string][]string)

func InitText() {
	remindText["english"] = []string{
		"After working for so long, it's time to give yourself a brief break. Have a drink of water and let your body and mind relax.                                          ",
		"Prolonged focus on work can lead to physical and mental fatigue. Now, let's momentarily shift our gaze, look around, and take a few deep breaths of fresh air.        ",
		"Rest is necessary to go the extra mile. Now, let's take a break, gaze out the window at the scenery, and give our minds a chance to unwind.                           ",
		"Prolonged focus can reduce efficiency. Why not stop now, do some simple stretches, and give your body a moment of relief?                                             ",
		"Work is important, but we must also learn to relax at the appropriate times. Now, put down your work, close your eyes, and take a few deep breaths.                   ",
		"A brief rest can help us regain our enthusiasm and focus for work. Now, stand up, take a few steps, and look around your environment.                                 ",
		"Sustained attention eventually wears out, so let your mind rest for a moment. Come, join me in doing some deep breathing and relaxation exercises.                    ",
		"Occasionally stepping away from the workplace can help us re-evaluate our state of mind. Now, gaze out the window and relax your tense nerves and muscles.            ",
		"On the long road of work, we need timely breaks to drink some water, relax our eyes and minds, and regain our focus.                                                  ",
		"Work is important, but we cannot ignore our own needs. Now, let's take a break, breathe deeply, and rejuvenate our bodies and minds in preparation for the work ahead.",
	}
	remindText["chinese"] = []string{
		"工作了这么久,是时候给自己一个短暂的休息了。来喝口水,让身体和大脑都好好放松一下吧。",
		"长时间专注工作容易导致身心疲惫,现在让我们暂时移开视线,环顾四周,深呼吸几口新鲜空气。",
		"休息是为了走更长远的路。现在就让我们稍作歇息,远眺窗外的风景,给心灵一个放松的机会。",
		"专注太久会降低效率,不如现在就停下来,做几个简单的伸展运动,让身体得到片刻的舒缓。",
		"工作固然重要,但我们也要学会在恰当的时候放松自己。现在就放下手头的工作,闭上眼睛深呼吸几次。",
		"一个短暂的休息能够让我们重拾工作的热情和专注力。现在就站起来,走几步,环顾四周的环境。",
		"持续的注意力总有耗尽的时候,此刻不妨让大脑小憩片刻。来,和我一起做几个深呼吸放松运动吧。",
		"偶尔远离工作岗位,能让我们重新审视自己的状态。现在就远眺窗外,放松你紧绷的神经和肌肉。",
		"在漫长的工作路上,我们需要适时的休憩,来喝口清水,放松一下你的眼睛和大脑,重拾专注力。",
		"工作固然重要,但我们也不能忽视自身的需求。现在就让我们稍作歇息,深呼吸,舒展身心,为接下来的工作做好准备。 都翻译成英文",
	}
}
func GetRandomRemindText() string {
	var textArray = remindText[config.GlobalConfig.Language.Language]
	rand.Seed(time.Now().Unix())
	return textArray[rand.Intn(len(textArray))]
}
